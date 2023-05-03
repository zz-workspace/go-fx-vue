<script lang="ts" setup>
import { LazyTableList, LazyWelcome, LazyTreeView, LazyBrowserTabs, LazyAppLogo, LazyEndpointQuery } from '#components'
import { ITab, ITable, IEndpoint } from '~/lib'
import { useTableStore } from '~/stores/table'

const { $listen } = useNuxtApp()
const browserTabsRef = ref<{
    newTab: (tab: ITab) => void
    activeTab: (tab: ITab) => void
}>()

const tableStore = useTableStore()
tableStore.setProjectId(1)

const tabs = shallowRef<ITab[]>([
    {
        label: 'Welcome',
        key: 'welcome',
        favico: () => {
            return iconMap.welcome
        },
        component: LazyWelcome,
        props: {
            key: 'home::welcome',
            name: 'google',
        },
    },
])

const tab = ref<string>(tabs.value[0].key)


$listen("sidebar:table:selected", async (table: ITable) => {
    const { fetchViews, fetchColumns } = tableStore;
    const columns = await fetchColumns(table.name)

    const views = await fetchViews(table.name)
    views?.map((view: any) => {
        for (const key in view) {
            const type = columns.find(column => column.name === key)?.type
            if (type === UITypes.JSON) {
                view[key] = decoded(view[key])
            }
        }
        return view
    })

    const key = 'table::' + table.name.toString()
    const existsTab = tabs.value.find(tab => tab.key === key)
    if (!existsTab) {
        const newTab: ITab = {
            component: LazyTableList,
            key: key,
            favico: () => {
                return iconMap.table
            },
            label: table.name,
            props: {
                tableName: table.name,
                columns: columns || [],
                rows: views || [],
            },
        }
        browserTabsRef.value?.newTab(newTab)
    } else {
        browserTabsRef.value?.activeTab(existsTab)
    }
})

$listen("sidebar:endpoint:selected", (endpoint: IEndpoint) => {
    const key = 'endpoint::' + endpoint.id.toString()
    const existsTab = tabs.value.find(tab => tab.key === key)
    if (!existsTab) {
        const newTab: ITab = {
            component: LazyEndpointQuery,
            key: key,
            favico: () => {
                return iconMap.api
            },
            label: endpoint.name,
            props: {
                tabKey: key,
                endpoint: endpoint,
            },
        }
        browserTabsRef.value?.newTab(newTab)
    } else {
        browserTabsRef.value?.activeTab(existsTab)
    }
})

const sidebarWidth = 280


</script>

<template>
    <div class="h-full">
        <div class="bg-primary-500" style="display: flex;">
            <div :style="{
                width: `${sidebarWidth}px`
            }">
                <LazyAppLogo />
            </div>
            <div :style="{ width: `calc(100% - ${sidebarWidth}px` }">
                <div class="bg-primary-500 pl-4">
                    <LazyBrowserTabs :tabs="tabs" v-model="tab" ref="browserTabsRef" />
                </div>
            </div>
        </div>
        <div style="height: calc(100vh - 50px);" class="flex">
            <div class="overflow-y-scroll" :style="{
                    height: '100%',
                    borderRight: '1px solid rgba(0, 0, 0, 0.12)',
                    width: sidebarWidth + 'px'
                }">
                <LazyTreeView />
            </div>
            <div class="flex-1 overflow-hidden">
                <div class="px-4">
                    <a-tabs v-model:activeKey="tab">
                        <a-tab-pane v-for="tabItem in tabs" :key="tabItem.key" :label="tabItem.label" :class="tabItem.key">
                            <component :is="tabItem.component" v-bind="tabItem.props" />
                        </a-tab-pane>
                    </a-tabs>
                </div>
            </div>
        </div>
    </div>
</template>
