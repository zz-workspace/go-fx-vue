<script setup lang="ts">
import { ITableColumn } from '~/lib'


const props = defineProps<{
    column: ITableColumn,
    editable: boolean
}>()
const emit = defineEmits(['save', 'delete'])
const column = toRef(props, 'column')

const editColumnDropdown = ref(false)
const formEdit = ref(false)

const showMenu = () => {
    editColumnDropdown.value = true
}

const handleRowSelect = () => {
    formEdit.value = true
}

const save = (newColumn: ITableColumn) => {
    emit('save', newColumn)
    formEdit.value = false;
}

</script>

<template>
    <div class="h-full select-none">
        <a-dropdown :trigger="['click']">
            <div @click="showMenu" class="h-full px-2">
                <div class="flex items-center justify-between h-full overflow-hidden">
                    <div class="flex items-center">
                        <LazySmartsheetHeaderCellIcon :column="column" />
                        <span class="ml-1">{{ column.name }}</span>
                    </div>
                    <component v-if="editable" :is="iconMap.arrowDown" />
                </div>
            </div>
            <template #overlay>
                <LazySmartsheetHeaderMenu v-if="editable" v-model="editColumnDropdown" @row-select="handleRowSelect" @delete="emit('delete')" />
            </template>

        </a-dropdown>
        <a-dropdown destroyPopupOnHide v-model:visible="formEdit" :trigger="['click']" placement="bottomRight">
            <div />
            <template #overlay>
                <LazySmartsheetHeaderFormAddOrEdit :column="column" @save="save" />
            </template>
        </a-dropdown>
    </div>
</template>