<script setup lang="ts">


const props = defineProps<{
    modelValue?: number | null
}>()

const vModel = useVModel(props, 'modelValue')

console.log(vModel.value)


</script>

<template>
    <span>{{ vModel }}</span>
</template>