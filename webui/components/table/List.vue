<script setup lang="ts">
import { ITableColumn, TableRowType } from '~/lib';


const props = defineProps<{
    rows: TableRowType[]
    columns: ITableColumn[]
    tableName: string
}>()

console.log(props.columns)

</script>

<template>
    <div>
        <LazySmartsheetGrid :rows="rows" :columns="columns" :tableName="tableName" />
    </div>
</template>