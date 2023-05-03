<script setup lang="ts">
import { ITableColumn } from '~/lib';
import { uiTypes } from '#imports'
const props = defineProps<{
    column?: ITableColumn
}>()

const emit = defineEmits(['save'])

const col = reactive({
    name: props.column?.name || '',
    type: props.column?.type || '',
    options: [...props.column?.options || []],
})

const uiTypesOptions = computed(() => uiTypes)

const save = () => {
    emit('save', col)
}

</script>

<template>
    <div class="w-[400px] bg-gray-50 shadow p-4 overflow-auto border">
        <a-form layout="vertical">
            <a-form-item label="Column Name" required>
                <a-input placeholder="input placeholder" v-model:value="col.name" />
            </a-form-item>
            <a-form-item label="Column Type" required>
                <a-select v-model:value="col.type">
                    <a-select-option v-for="opt of uiTypesOptions" :key="opt.name" :value="opt.name">
                        <div class="flex gap-1 items-center">
                            <component :is="opt.icon" class="text-grey" />
                            {{ opt.name }}
                        </div>
                    </a-select-option>
                </a-select>
            </a-form-item>
            <LazySmartsheetColumnSelectOptions v-model="col.options"
                v-if="col.type === UITypes.SingleSelect || col.type === UITypes.MultipleSelect" />

            <a-form-item>
                <div class="flex justify-end">
                    <a-button>Cancel</a-button>
                    <a-button type="primary" class="ml-2" @click="() => save()">Save</a-button>
                </div>
            </a-form-item>
        </a-form>
    </div>
</template>