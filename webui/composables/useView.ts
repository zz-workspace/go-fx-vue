import { message } from "ant-design-vue"
import { ITableColumn, TableRowType } from "~/lib"
import { useTableStore } from "~/stores/table"
import dayjs from 'dayjs'

export const useView = (tableName: string) => {
    const views = ref<TableRowType[]>([])
    const tableStore = useTableStore()

    const formatView = (columns: ITableColumn[], view: TableRowType) => {
        for (const key in view) {
            const type = columns.find(column => column.name === key)?.type
            if (type === UITypes.JSON) {
                view[key] = decoded(view[key])
            }
            if (type === UITypes.DateTime) {
                view[key] = dayjs(view[key]).format('YYYY-MM-DD HH:mm:ss')
            }
        }

        return view;
    }

    const loadViews = async (columns: ITableColumn[]) => {
        const $views = await tableStore.fetchViews(tableName)
        $views?.map((view: any) => {
           return formatView(columns, view)
        })
        views.value = $views
    }

    const addView = async (columns: ITableColumn[],) => {
        let view = await tableStore.addView(tableName)
        view =  formatView(columns, view)
        if (views.value) {
            views.value?.push(view)
        } else {
            views.value = [view]
        }
        message.success('Row Added');
    }

    const updateView = async (viewId: number | string, key: string, value: any) => {
        await tableStore.updateView(tableName, viewId, {
            key: key,
            value: value,
        })

        message.success('Row Updated')
    }

    const deleteView = async (viewIndex: number) => {
        await tableStore.deleteView(tableName, views.value![viewIndex].id)
        views.value?.splice(viewIndex, 1)
        message.success('Row Deleted')
    }

    const clearViewCell = async (columnName: string, viewIndex: number) => {
        views.value![viewIndex][columnName] = null
    }

    return {
        views,
        loadViews,
        addView,
        updateView,
        deleteView,

        clearViewCell,
    }
}