<script setup lang="ts">
import dayjs from 'dayjs'



const props = defineProps<{
    readonly?: boolean,
    modelValue?: string | null
}>()


const emit = defineEmits(['update:modelValue', 'change'])

const vModel = useVModel(props, 'modelValue', emit)

const date = ref(dayjs(vModel.value))

watch(date, () => {
    if (date.value) {
        emit('update:modelValue', date.value.format('YYYY-MM-DD HH:mm:ss'))
    } else {
        emit('update:modelValue', '')
    }
    emit('change', date.value.format('YYYY-MM-DD HH:mm:ss'))
})

</script>

<template>
    <a-date-picker show-time v-if="!readonly" v-model:value="date" :bordered="false" />
    <span v-else>{{ vModel }}</span>
</template>