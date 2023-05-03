<script setup lang="ts">
import { HTMLElementEvent } from '~/lib'



const props = defineProps<{
    readonly?: boolean,
    modelValue?: boolean | string | null
}>()

const emit = defineEmits(['change', 'update:modelValue'])

const vModel = useVModel(props, 'modelValue')

const handleChange = (event: Event) => {
    emit('change', (event as HTMLElementEvent<HTMLInputElement>).target.checked)
}

</script>

<template>
    <div class="flex justify-center items-center w-full">
        <a-checkbox v-model:checked="vModel" @change="handleChange"></a-checkbox>
    </div>
</template>

<style lang="scss" scoped>
.cell-text {
    outline: none;
    border: none;
    padding: 0 5px;
    width: 100%;
}
</style>