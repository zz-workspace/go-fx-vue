import { defineStore } from "pinia"
import { ITable, ITableColumn, TableRowType } from "~/lib";
import lodash from 'lodash'
import { DataTypes } from "~/utils/columnUtil";
import { PgUi } from "~/lib/sqlUi/PgUi";

export const useTableStore = defineStore('table', () => {
    const columns = ref([])
    const tables = ref([])
    const views = ref([])
    const { $api } = useNuxtApp()
    const projectId = ref<number | string>()
    const uri = computed(() => `/projects/${projectId.value}`)

    const fetchTables = async () => {
        const response = await $api.get(`${uri.value}/tables`);
        tables.value = response.data;
        return response.data;

    }

    const createTable = async (table: ITable) => {
        await $api.post(`${uri.value}/tables`, table);
        await fetchTables()
    }

    const fetchColumns = async (tableName: string) => {
        const response = await $api.get(`${uri.value}/tables/${tableName}/columns`);
        const columns: ITableColumn[] = (response.data as any[])?.map(col => {
            let columnType = PgUi.getUiTypeFromDataType(col.data_type)
            return {
                name: col.name,
                options: col.options,
                type: columnType,
            }
        })

        return lodash.sortBy(columns, (column) => column.name === 'id' ? 0 : 1)
    }

    const updateColumn = async (tableName: string, data: { old_column: string; new_column: string; data_type: string }) => {
        const response = await $api.put(`${uri.value}/tables/${tableName}/columns`, data);
        return response.data;
    }

    const insertColumn = async (tableName: string, column: { name: string; type: DataTypes }) => {
        const response = await $api.post(`${uri.value}/tables/${tableName}/columns`, column);
        return response.data;
    }

    const deleteColumn = async (tableName: string, column: { name: string }) => {
        const response = await $api.delete(`${uri.value}/tables/${tableName}/columns/${column.name}`);
        return response.data;
    }

    const fetchViews = async (tableName: string) => {
        const response = await $api.get(`${uri.value}/tables/${tableName}/views`);
        return response.data;
    }

    const setProjectId = ($projectId: number | string) => {
        projectId.value = $projectId
    }

    const addView = async (tableName: string) => {
        const response = await $api.post(`${uri.value}/tables/${tableName}/views`);
        return response.data
    }

    const updateView = async (tableName: string, rowId: any, row: { key: string; value: any }) => {
        const response = await $api.put(`${uri.value}/tables/${tableName}/views/${rowId}`, row);
        return response.data
    }

    const deleteView = async (tableName: string, viewId: any) => {
        const response = await $api.delete(`${uri.value}/tables/${tableName}/views/${viewId}`);
        return response.data
    }

    return {
        tables,
        columns,
        views,
        createTable,
        fetchTables,
        setProjectId,
        fetchColumns,
        updateColumn,
        insertColumn,
        deleteColumn,
        fetchViews,
        addView,
        updateView,
        deleteView,
    }
})