<script lang="ts" setup>
import { Modal } from 'ant-design-vue';
import { resolve } from 'path';
import { h } from 'vue';
import { ITable, IEndpoint } from '~/lib'
import { useTableStore } from '~/stores/table';
import { TableList } from '~~/.nuxt/components';

const endpoints: any[] = [
    {
        id: 1,
        name: 'users',
        children: [
            {
                id: 1,
                name: 'getUser',
                method: 'GET'
            },
            {
                id: 2,
                name: 'updateUser',
                method: 'PUT'
            },
            {
                id: 3,
                name: 'deleteUser',
                method: 'DELETE'
            }
        ]
    },

]


const openTableCreateDialog = ref(false)

const { $event } = useNuxtApp()
const selectedTable = (table: ITable) => {
    $event('sidebar:table:selected', table)
}

const tableStore = useTableStore()

const tables = computed(() => tableStore.tables)


onMounted(async () => {
    await tableStore.fetchTables()
})

const selectedEndpoint = (endpoint: IEndpoint) => {
    $event('sidebar:endpoint:selected', endpoint)
}



</script>

<template>
    <div class="py-4">
        <div>
            <span class="text-base px-4">TABLES</span>
            <div class="mt-2">
                <div class="flex items-center py-2 px-4 cursor-pointer hover:text-primary-500/100"
                    @click="openTableCreateDialog = true">
                    <component :is="iconMap.plus" />
                    <span class="ml-2">Add New Table</span>
                </div>
                <div class="flex items-center py-2 px-4 cursor-pointer" v-for="table in tables" :key="table.id" :class="{
                        // 'bg-gray-100': tableSelected.id === table.id
                    }" @click="selectedTable(table)">
                    <component :is="iconMap.table" class="text-primary-500" />
                    <span class="ml-2">{{ table.name }}</span>
                </div>
            </div>
        </div>
        <a-divider />
        <div class="mt-4">
            <span class="text-base px-4">API</span>
            <div class="mt-2">
                <div class="flex items-center py-2 px-4 cursor-pointer hover:text-primary-500/100">
                    <component :is="iconMap.plus" />
                    <span class="ml-2">Add New API</span>
                </div>
                <div class="px-4 cursor-pointer" v-for="api in endpoints" :key="api.id">
                    <div class="flex items-center py-2">
                        <component :is="iconMap.api" class="text-primary-500" />
                        <span class="   ml-2">{{ api.name }}</span>
                    </div>
                    <div v-for="endpoint in api.children" class="pl-6">
                        <div class="flex items-center py-2" @click="() => selectedEndpoint(endpoint)">
                            <a-tag v-if="endpoint.method === 'GET'" color="green">{{ endpoint.method }}</a-tag>
                            <a-tag v-if="endpoint.method === 'POST'" color="blue">{{ endpoint.method }}</a-tag>
                            <a-tag v-if="endpoint.method === 'PUT'" color="orange">{{ endpoint.method }}</a-tag>
                            <a-tag v-if="endpoint.method === 'DELETE'" color="red">{{ endpoint.method }}</a-tag>

                            <span class="ml-1">{{ endpoint.name }}</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <a-divider />
        <div class="mt-4">
            <span class="text-base px-4">Functions</span>
        </div>
        <a-divider />
        <div class="mt-4">
            <span class="text-base px-4">Tasks</span>
        </div>
        <LazyDialogTableCreate v-model="openTableCreateDialog" />
    </div>
</template>