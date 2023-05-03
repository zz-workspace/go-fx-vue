<script setup lang="ts">
import { vOnClickOutside } from '@vueuse/components'
import { ITableColumn } from '~/lib'
const props = defineProps({
    modelValue: {
        type: [String, Number, Array, Boolean] as PropType<any>
    },
    column: {
        type: Object as PropType<ITableColumn>
    }
})
const emit = defineEmits(['update:modelValue', 'change'])

const vModel = useVModel(props, 'modelValue', emit)

// watch(vModel, () => {
//     emit('change', vModel.value)
// })

const handleChange = (value: any) => {
    emit('change', value)
}

const readonly = ref(props.column?.type !== UITypes.ID)
const active = ref(false)

const handleClickOutSide = () => {
    readonly.value = true
    setReadonly(true)
    active.value = false
}

const handeDoubleClick = () => {
    if (props.column?.name === 'id') {
        return;
    }
    readonly.value = false
}

const setReadonly = (isReadonly: boolean) => {
    readonly.value = isReadonly
}

const setActive = () => {
    if (props.column?.name === 'id') {
        return;
    }
    active.value = true
}


onKeyStroke(['Enter'], () => {
    setReadonly(true)
    active.value = false
});

</script>

<template>
    <div class="grid-cell h-full px-4" @click="setActive" :class="{ active: active }" @dblclick="handeDoubleClick"
        v-on-click-outside="[handleClickOutSide, {
                ignore: ['.ant-select-dropdown', '.ant-picker-dropdown'],
            }]">
        <LazyCellID v-model="vModel" v-if="column?.type === UITypes.ID" @change="handleChange" />
        <LazyCellText v-model="vModel" v-if="column?.type === UITypes.Text" :readonly="readonly" @change="handleChange" />
        <LazyCellSingleSelect :options="props.column?.options" v-else-if="column?.type === UITypes.SingleSelect"
            @change="handleChange" v-model="vModel" :readonly="readonly" />
        <LazyCellMultipleSelect :options="props.column?.options" v-else-if="column?.type === UITypes.MultipleSelect"
            @change="handleChange" v-model="vModel" :readonly="readonly" />
        <LazyCellCheckbox v-else-if="column?.type === UITypes.Checkbox" v-model="vModel" :readonly="readonly"
            @change="handleChange" />
        <LazyCellEmail v-else-if="column?.type === UITypes.Email" v-model="vModel" :readonly="readonly"
            @change="handleChange" />
        <LazyCellDateTime v-else-if="column?.type === UITypes.DateTime" v-model="vModel" :readonly="readonly"
            @change="handleChange" />
        <LazyCellPassword v-else-if="column?.type === UITypes.Password" v-model="vModel" :readonly="readonly"
            @change="handleChange" />
        <LazyCellNumber v-else-if="column?.type === UITypes.Number" v-model="vModel" :readonly="readonly"
            @change="handleChange" />
        <LazyCellOther v-else="column?.type === UITypes.Other" v-model="vModel" :readonly="readonly"
            @change="handleChange" />
    </div>
</template>

<style lang="scss" scoped>
.grid-cell {
    display: flex;
    align-items: center;
    box-sizing: content-box;

    &.active {
        @apply border-primary-500 border-1;
    }
}
</style>