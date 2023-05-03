<script setup lang="ts">

const props = defineProps({
    modelValue: {
        type: Boolean as PropType<boolean>
    }
})
const { $event } = useNuxtApp()

const emit = defineEmits(['rowSelect', 'update:modelValue', 'delete'])


</script>

<template>
    <a-menu>
        <a-menu-item @click="emit('rowSelect')">
            <div class="flex items-center">
                <component :is="iconMap.edit" class="text-primary" /> <span class="ml-2">Edit</span>
            </div>
        </a-menu-item>
        <a-menu-item>
            <div class="flex items-center">
                <component :is="iconMap.sortAsc" class="text-primary" /> <span class="ml-2">Sort Ascending</span>
            </div>
        </a-menu-item>
        <a-menu-item>
            <div class="flex items-center">
                <component :is="iconMap.sortDesc" class="text-primary" /> <span class="ml-2">Sort Descending</span>
            </div>
        </a-menu-item>
        <a-menu-item>
            <div class="flex items-center">
                <component :is="iconMap.eye" class="text-primary" /> <span class="ml-2">Hide Field</span>
            </div>
        </a-menu-item>
        <a-menu-item>
            <div class="flex items-center" @click="emit('delete')">
                <component :is="iconMap.delete" class="text-red" /> <span class="ml-2">Delete</span>
            </div>
        </a-menu-item>
    </a-menu>
</template>