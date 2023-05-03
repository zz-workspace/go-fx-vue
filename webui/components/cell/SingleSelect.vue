<script setup lang="ts">

const props = defineProps<{
    readonly?: boolean,
    modelValue?: string | null
    options: string[]
}>()
const emit = defineEmits(['change'])
const cellTextRef = ref()

const vModel = useVModel(props, 'modelValue')

watch(cellTextRef, () => {
    if (cellTextRef.value) {
        cellTextRef.value.focus()
    }

})

const singleSelectRef = ref()

defineExpose({
    singleSelectRef
})

</script>

<template>
    <a-select v-if="!readonly" v-model:value="vModel" class="cell-single-select" @change="emit('change')">
        <a-select-option :value="option" v-for="option in props.options" :key="option">{{ option }}</a-select-option>
    </a-select>

    <span v-else>{{ vModel }}</span>
</template>

<style lang="scss" scoped>
.cell-single-select {
    @apply w-full;

    :deep(.ant-select-selector) {
        @apply border-none h-full !shadow-none;

        .ant-select-selection-item {
            @apply flex items-center;
        }
    }


}
</style>