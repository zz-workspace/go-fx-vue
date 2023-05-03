<script setup lang="ts">

const props = defineProps({
    modelValue: {
        type: Array as PropType<string[]>
    }
})

const emit = defineEmits(['update:modelValue'])


const selectOptions = reactive<string[]>(props.modelValue || [])

const addOption = () => {
    selectOptions.push('')
}

const removeOption = (index: number) => {
    selectOptions.splice(index, 1)
}


watch(selectOptions, () => {
    emit('update:modelValue', selectOptions.filter(opt => opt))
})


</script>

<template>
    <div class="w-full">
        <a-space direction="vertical" class="w-full overflow-auto max-h-xs">
            <div v-for="(opt, index) in selectOptions" class="flex items-center">
                <a-input v-model:value="selectOptions[index]" />
                <component :is="iconMap.close" class="cursor-pointer ml-2" @click="() => removeOption(index)" />
            </div>

        </a-space>

        <a-button type="dashed" class="w-full flex items-center my-4" @click="addOption">
            <div class="flex">
                <component :is="iconMap.plus" />
                <span class="flex-auto">Add option</span>
            </div>
        </a-button>
    </div>
</template>