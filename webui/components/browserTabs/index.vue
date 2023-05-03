<script lang="ts" setup>
import Vue3TabsChrome from 'vue3-tabs-chrome'
import 'vue3-tabs-chrome/dist/vue3-tabs-chrome.css'
import { ITab } from '~/lib';
const props = defineProps<{
    tabs: ITab[];
    modelValue: any;
}>()

const emit = defineEmits(['update:modelValue'])

const browserTabsRef = ref()

const { modelValue } = props;

const tabSelected = ref(modelValue)

const tabs = ref(props.tabs)

watch(tabSelected, () => {
    emit('update:modelValue', tabSelected.value)
})

onMounted(() => {
    console.log(browserTabsRef.value)
})

const newTab = (tab: ITab) => {
    browserTabsRef.value.addTab(tab)
    tabSelected.value = tab.key;
}


const activeTab = (tab: ITab) => {
    tabSelected.value = tab.key;
}

defineExpose({
    newTab,
    activeTab,
})
</script>
<template>
    <vue3-tabs-chrome ref="browserTabsRef" class="browser-tabs" :tabs="tabs" v-model="tabSelected" />
</template>

<style lang="scss" scoped>
.browser-tabs {
    @apply !bg-primary-500 padding-top: -10px;
    :deep(.tabs-background-divider) {
        @apply d-none;
    }

    :deep(.tabs-content) {
        @apply h-10;

        .tabs-item {
            // width: 240px !important;

            .tabs-label {
                @apply text-white;
            }

            .tabs-background-content {
                @apply bg-primary-500;
            }

            .tabs-close-icon {
                @apply stroke-gray-500;
            }

            .tabs-main {
                margin: 0 20px;
            }

            .tabs-favico {
                span {
                    @apply text-white;
                }
            }

            &.active {
                .tabs-label {
                    @apply text-primary-500;
                }

                .tabs-background-content {
                    @apply bg-white;
                }

                .tabs-favico {
                    span {
                        @apply text-primary-500;
                    }
                }

            }

            .tabs-background-divider {
                display: none;
            }

            &:hover {
                .tabs-background-content {
                    @apply bg-white;
                }

                .tabs-label {
                    @apply text-primary-500;
                }

                .tabs-favico {
                    span {
                        @apply text-primary-500;
                    }
                }
            }
        }
    }



}
</style>