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
    const value = (event as HTMLElementEvent<HTMLInputElement>).target.value
    emit('update:modelValue', value)
    emit('change', value)
}
</script>

<template>
    <input ref="cellTextRef" autofocus v-if="!readonly" class="cell-text" @change="handleChange" :value="vModel"
        @blur="emit('blur')" />
    <a v-else :href="'mailto:' + vModel">{{ vModel }}</a>
</template>

<style lang="scss" scoped>
.cell-text {
    outline: none;
    border: none;
    padding: 0 5px;
    width: 100%;
}
</style>