<script setup lang="ts">
import { Modal } from 'ant-design-vue';
import { ITableColumn } from '~/lib';
const props = defineProps<{
    tableName: string
}>()

const { columns, loadColumns, addColumn, updateColumn, deleteColumn, insertColumn } = useColumn(props.tableName)
const { views, loadViews, addView, updateView, deleteView, clearViewCell } = useView(props.tableName)

onMounted(async () => {
    await loadColumns()
    await loadViews(columns.value)
})
const contextMenuTarget = reactive<{
    rowIndex: number | null;
    colIndex: number | null;
    column?: ITableColumn
}>({
    rowIndex: null,
    colIndex: null,
})

const showAddColumn = ref(false)
const handleAddColumn = (newColumn: ITableColumn) => {
    insertColumn(newColumn)
    showAddColumn.value = false;
}

const handleDeleteView = async () => {
    if (typeof contextMenuTarget.rowIndex === 'number') {
        deleteView(contextMenuTarget.rowIndex)
    }
}

const handleDeleteColumn = (colIndex: number) => {
    Modal.confirm({
        title: 'Do you want to delete these items?',
        icon: iconMap.account,
        content: 'When clicked the OK button, this dialog will be closed after 1 second',
        onOk() {
            deleteColumn(colIndex)
        },
        onCancel() { },
    })
}

const handleClearViewCell = () => {
    if (typeof contextMenuTarget.colIndex === 'number' && typeof contextMenuTarget.rowIndex === 'number') {
        const col = columns.value[contextMenuTarget.colIndex]
        clearViewCell(col.name, contextMenuTarget.rowIndex)
    }
}

const handleUpdateColumn = async (column: ITableColumn, newColumn: ITableColumn) => {
    const columns = await updateColumn(column, newColumn)
    loadViews(columns)
}

const showContextMenu = (rowIndex: number, columnIndex: number, column: ITableColumn) => {
    contextMenuTarget.rowIndex = rowIndex;
    contextMenuTarget.colIndex = columnIndex;
    contextMenuTarget.column = column
}

</script>

<template>
    <div class="grid-wrapper">
        <a-dropdown :trigger="['contextmenu']">
            <table>
                <thead>
                    <tr>

                        <th class="cursor-pointer" v-for="(column, index) in columns" :key="index" :class="{
                                ['grid-cell-' + index]: true,
                            }">
                            <LazySmartsheetHeaderCell @save="(newColumn: ITableColumn) => handleUpdateColumn(column, newColumn)"
                                @delete="handleDeleteColumn(index)" :editable="column.name !== 'id'" :column="column" />
                        </th>

                        <th class="btn-add-column" @click="showAddColumn = true">
                            <div class="h-full flex items-center justify-center cursor-pointer">
                                <component :is="iconMap.plus" />
                            </div>
                            <a-dropdown destroyPopupOnHide v-model:visible="showAddColumn" :trigger="['click']"
                                placement="bottomRight">
                                <div />
                                <template #overlay>
                                    <LazySmartsheetHeaderFormAddOrEdit @save="handleAddColumn" />
                                </template>
                            </a-dropdown>
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(view, viewIndex) in views" :key="view.id">
                        <td v-for="(column, columnIndex) in columns" :key="column.name"
                            @contextmenu="() => showContextMenu(viewIndex, columnIndex, column)">
                            <LazySmartsheetCell :column="column" v-model="view[column.name]"
                                @change="(value: any) => updateView(view.id as number, column.name, value)">
                            </LazySmartsheetCell>
                        </td>

                    </tr>
                    <tr>
                        <td :colspan="columns.length" class="cursor-pointer" @click="addView(columns)">
                            <div class="text-gray-500 flex items-center">
                                <component :is="iconMap.plus" />
                                <span class="ml-2">Add new row</span>
                            </div>
                        </td>
                    </tr>
                </tbody>
            </table>
            <template #overlay>
                <a-menu>
                    <a-menu-item @click="handleDeleteView">
                        <div class="flex items-center">
                            <span>Delete Row</span>
                        </div>
                    </a-menu-item>
                    <a-menu-item @click="handleClearViewCell" v-if="contextMenuTarget.column?.name !== 'id'">
                        <div class="flex items-center">
                            <span>Clear Cell</span>
                        </div>
                    </a-menu-item>

                </a-menu>
            </template>
        </a-dropdown>

    </div>
</template>

<style lang="scss" scoped>
.grid-wrapper {
    overflow: auto;
    height: 800px;
    width: 100%;

    table {
        border-spacing: 0;
        background-color: #fff;

        thead {
            tr {
                th {
                    width: 200px;
                    min-width: 200px;
                    max-width: 200px;
                    height: 41px;
                    background-color: rgba(243, 244, 246, 1);
                    border-right: 1px solid rgba(229, 231, 235, 1);
                    border-bottom: 1px solid rgba(229, 231, 235, 1);
                    color: rgba(107, 114, 128, 1);
                    font-weight: normal;
                    text-align: left;
                    position: sticky;
                    top: 0;

                    &.btn-add-column {
                        // display: flex;
                        // justify-content: center;
                        // align-items: center;
                        // width: 64px;
                        // min-width: 64px;
                        // max-width: 64px;
                        @apply w-10 min-w-10 max-w-10;
                    }
                }
            }
        }

        tbody {
            tr {
                td {
                    width: 200px;
                    min-width: 200px;
                    max-width: 200px;
                    height: 41px;
                    background-color: #fff;
                    border-right: 1px solid rgba(229, 231, 235, 1);
                    border-bottom: 1px solid rgba(229, 231, 235, 1);
                    color: #000;
                    font-weight: normal;
                    overflow: hidden;


                }
            }
        }
    }
}
</style>