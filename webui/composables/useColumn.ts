import { message } from "ant-design-vue"
import { ITableColumn } from "~/lib"
import { PgUi } from "~/lib/sqlUi/PgUi"
import { useTableStore } from "~/stores/table"

export const useColumn = (tableName: string) => {
    const columns = ref<ITableColumn[]>([])
    const tableStore = useTableStore()

    const loadColumns = async () => {
        columns.value = await tableStore.fetchColumns(tableName)
    }

    const addColumn = async (column: ITableColumn) => {
        columns.value.push(column)
    }

    const updateColumn = async (oldColumn: ITableColumn, newColumn: ITableColumn) => {
        await tableStore.updateColumn(tableName, {
            old_column: oldColumn.name,
            new_column: newColumn.name,
            data_type: newColumn.type,
        })
        const colIndex = columns.value.findIndex(column => column.name === oldColumn.name)
        columns.value[colIndex].name = newColumn.name
        message.success('Column Updated')
        return columns.value;
    }

    const insertColumn = async (column: ITableColumn) => {
        await tableStore.insertColumn(tableName, {
            name: column.name,
            type: PgUi.getDataTypeFromUiType(column.type),
        })
        await loadColumns()
        message.success('Column Inserted')
    }

    const deleteColumn = async (columnIndex: number) => {
        const column = columns.value[columnIndex]
        await tableStore.deleteColumn(tableName, column)
        columns.value.splice(columnIndex, 1)
        message.success('Column Deleted')
    }

    return {
        columns,
        loadColumns,
        updateColumn,
        addColumn,
        deleteColumn,
        insertColumn,
    }
}