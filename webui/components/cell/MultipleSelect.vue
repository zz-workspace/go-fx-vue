<script setup lang="ts">

const props = defineProps<{
    readonly?: boolean,
    modelValue?: string[] | null
    options: string[]
}>()
const emit = defineEmits(['change'])

const vModel = useVModel(props, 'modelValue')

const handleChange = (e) => {
    emit('change', e.target.value)
}

</script>

<template>
    <a-select mode="multiple" v-if="!readonly" v-model:value="vModel" class="cell-multiple-select" @change="handleChange">
        <a-select-option :value="option" v-for="option in props.options" :key="option">{{ option }}</a-select-option>
    </a-select>
    <div v-else class="flex">
        <a-tag v-for="opt in vModel" :key="opt">{{ opt }}</a-tag>
    </div>
</template>

<style lang="scss" scoped>
.cell-multiple-select {
    @apply w-full;

    :deep(.ant-select-selector) {
        @apply border-none h-full !shadow-none;

        .ant-select-selection-item {
            @apply flex items-center;
        }
    }


}
</style>