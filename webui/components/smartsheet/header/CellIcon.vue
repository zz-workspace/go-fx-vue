<script setup lang="ts">
import { ITableColumn } from '~/lib';

const props = defineProps<{
    column: ITableColumn
}>()
const type = computed(() => props.column.type)
const name = computed(() => props.column.name)
</script>

<template>
    <div v-if="name === 'id'">#</div>
    <component v-else-if="type === UITypes.Text" :is="iconMap.text" />
    <component v-else-if="type === UITypes.Email" :is="iconMap.email" />
    <component v-else-if="type === UITypes.SingleSelect" :is="iconMap.singleSelect" />
    <component v-else-if="type === UITypes.MultipleSelect" :is="iconMap.multiSelect" />
    <component v-else-if="type === UITypes.Checkbox" :is="iconMap.check" />
    <component v-else-if="type === UITypes.DateTime" :is="iconMap.datetime" />
    <component v-else-if="type === UITypes.Password" :is="iconMap.passwordChange" />
    <component v-else-if="type === UITypes.Number" :is="iconMap.number" />
    <component v-else-if="type === UITypes.Other" :is="iconMap.other" />
    <component v-else-if="type === UITypes.JSON" :is="iconMap.json" />
</template>