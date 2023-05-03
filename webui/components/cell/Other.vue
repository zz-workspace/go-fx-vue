<script setup lang="ts">
import { HTMLElementEvent } from '~/lib';


const props = defineProps<{
    readonly?: boolean,
    modelValue?: string | null
}>()

const cellTextRef = ref()

const vModel = useVModel(props, 'modelValue')
const emit = defineEmits(['blur', 'update:modelValue', 'change'])

watch(cellTextRef, () => {
    if (cellTextRef.value) {
        cellTextRef.value.focus()
    }

})

const handleChange = (event: Event) => {
    emit('update:modelValue', (event as HTMLElementEvent<HTMLInputElement>).target.value)
        emit('change')
}
</script>

<template>
    <input ref="cellTextRef" autofocus v-if="!readonly" class="cell-text" :value="vModel" @change="handleChange"
        @blur="emit('blur')" />
    <span v-else>{{ vModel }}</span>
</template>

<style lang="scss" scoped>
.cell-text {
    outline: none;
    border: none;
    padding: 0 5px;
    width: 100%;
}
</style>