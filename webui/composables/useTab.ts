import { ITab } from "~/lib"

export const useTab = () => {
    const tabs = ref<ITab[]>([])

    const addTab = (tab: ITab) => {

    }

    return {
        tabs,
        addTab,
    }
}