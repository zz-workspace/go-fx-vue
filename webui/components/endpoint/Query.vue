<script setup lang="ts">
import { IEndpoint } from '~/lib';
const props = defineProps<{
    tabKey: string
}>()

console.log(props.tabKey)
</script>
<template>
    <div>Endpoint {{ props.tabKey }}</div>
</template>