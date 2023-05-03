package controller

import (
	"fmt"
	"net/http"

	"fast-api.io/internal/repository"
	"fast-api.io/internal/service"
	"fast-api.io/models"
	"fast-api.io/modules/http/response"
	"fast-api.io/modules/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/swarm"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type WorkspaceRequest struct {
	Name string `json:"name"`
}

type DeleteWorkspaceRequest struct {
	IDS []uint64 `json:"ids"`
}

type WorkspaceController struct {
	workspaceRepository repository.WorkspaceRepository
	dockerService       service.DockerService
}

func InitWorkspaceController(r *gin.RouterGroup, workspaceRepository *repository.WorkspaceRepository, dockerService *service.DockerService) {
	controller := &WorkspaceController{
		workspaceRepository: *workspaceRepository,
		dockerService:       *dockerService,
	}
	routes := r.Group("workspaces")
	routes.GET("", controller.WorkspaceList)
	routes.POST("", controller.CreateWorkspace)
	routes.DELETE("", controller.DeleteWorkspace)
}

func (e WorkspaceController) CreateWorkspace(ctx *gin.Context) {
	var body WorkspaceRequest
	var validate = validator.New()
	ctx.ShouldBindJSON(&body)
	err1 := validate.Struct(body)
	if err1 != nil {
		response.ValidationError(ctx, http.StatusBadRequest, err1)
		return
	}

	workspace := e.workspaceRepository.FindByName(body.Name)
	fmt.Println(workspace.ID)
	if workspace.ID != 0 {
		response.Error(ctx, http.StatusConflict, fmt.Errorf("Workspace already exists!"))
		return
	}

	// Create postgresQL service (only https)
	appName := body.Name
	// traefikEntryPointsLabel := fmt.Sprint("traefik.tcp.routers.", appName, ".entryPoints")
	traefikRuleLabel := fmt.Sprint("traefik.tcp.routers.", appName, ".rule")
	// traefikServiceLabel := fmt.Sprint("traefik.tcp.routers.", appName, ".service")
	traefikLoadBanlancerLabel := fmt.Sprint("traefik.tcp.services.", appName, ".loadbalancer.server.port")
	// traefikTLSLabel := fmt.Sprint("traefik.tcp.routers.", appName, ".tls.certresolver")
	traefikTLSLabel := fmt.Sprint("traefik.tcp.routers.", appName, ".tls")
	origin := fmt.Sprint(appName, ".db.docker.localhost")
	var host string = fmt.Sprint("HostSNI(`", origin, "`)")
	serviceId, err := e.dockerService.CLI.ServiceCreate(
		ctx,
		swarm.ServiceSpec{
			Mode: swarm.ServiceMode{
				Replicated: &swarm.ReplicatedService{Replicas: utils.GetIntPointer(1)},
			},
			Annotations: swarm.Annotations{
				Name: appName,
			},
			Networks: []swarm.NetworkAttachmentConfig{
				{
					Target: "mynetwork",
				},
			},
			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{

					Image: "postgres",
					Labels: map[string]string{
						"traefik.enable": "true",
						// traefikEntryPointsLabel: "websecure",
						traefikRuleLabel: host,

						// traefikServiceLabel:       "postgres",
						traefikLoadBanlancerLabel: "5432",
						traefikTLSLabel:           "true",
					},

					Mounts: []mount.Mount{
						{
							Source: appName,
							Type:   "volume",
							Target: "/var/lib/postgresql/data",
						},
					},
					Env: []string{
						"POSTGRES_DB=demo",
						"POSTGRES_USER=user",
						"POSTGRES_PASSWORD=abcd1234",
					},
				}}}, types.ServiceCreateOptions{})

	if err != nil {
		panic(err)
	}

	workspaceResponse := e.workspaceRepository.CreateWorkspace(&models.Workspace{
		Name:        body.Name,
		DBServiceID: serviceId.ID,
		DBOrigin:    origin,
	})
	response.JSON(ctx, http.StatusCreated, workspaceResponse)
}

func (e WorkspaceController) DeleteWorkspace(ctx *gin.Context) {
	var body DeleteWorkspaceRequest
	var validate = validator.New()
	ctx.ShouldBindJSON(&body)
	err1 := validate.Struct(body)
	if err1 != nil {
		response.ValidationError(ctx, http.StatusBadRequest, err1)
		return
	}
	for _, id := range body.IDS {
		workspace := e.workspaceRepository.FindByID(id)
		if workspace.ID != 0 {

			filters := filters.NewArgs()
			filters.Add("service", workspace.Name)
			// tasks, _ := e.dockerService.CLI.TaskList(ctx, types.TaskListOptions{
			// 	Filters: filters,
			// })
			e.dockerService.CLI.ServiceRemove(ctx, workspace.DBServiceID)
			// xxx := make(chan tasks)
			// for _, task := range tasks {
			// 	statusCh, errCh := e.dockerService.CLI.ContainerWait(ctx, task.Status.ContainerStatus.ContainerID, container.WaitConditionRemoved)
			// 	select {
			// 	case err := <-errCh:
			// 		if err != nil {
			// 			panic(err)
			// 		}
			// 	case <-statusCh:
			// 	}
			// }

			// err := e.dockerService.CLI.VolumeRemove(ctx, workspace.Name, true)
			// if err != nil {
			// 	panic(err)
			// }

		}
	}
	workspaces := e.workspaceRepository.DeleteManyWorkspaces(body.IDS)
	response.JSON(ctx, http.StatusOK, workspaces)
}

func (e WorkspaceController) WorkspaceList(ctx *gin.Context) {
	workspaces, _ := e.workspaceRepository.WorkspaceList()
	response.JSON(ctx, http.StatusOK, workspaces)
}
