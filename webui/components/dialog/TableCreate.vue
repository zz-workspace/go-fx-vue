<script setup lang="ts">
import { ITable } from '~/lib'
import { useTableStore } from '~/stores/table'

const props = defineProps<{
    modelValue: boolean
}>()
const emit = defineEmits(['update:modelValue'])
const dialogShow = useVModel(props, 'modelValue', emit)
const loading = ref(false)
const tableStore = useTableStore()
const table = reactive<ITable>({
    name: '',
})

const handleOk = async () => {
    loading.value = true
    await tableStore.createTable(table)
}
</script>

<template>
    <a-modal v-model:visible="dialogShow" title="Table Create" ok-text="Save" :confirmLoading="loading" @ok="handleOk">
        <a-form layout="vertical">
            <a-form-item label="Table Name" required placeholder="Enter Table Name">
                <a-input v-model:value="table.name" />
            </a-form-item>
        </a-form>
    </a-modal>
</template>