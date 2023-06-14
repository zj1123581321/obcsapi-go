import { ref, type Ref, onMounted } from 'vue'
import { defineStore } from 'pinia'
import { ObcsapiMentionGet } from "@/api/obcsapi"

class LocalSettingsClass {
    LoadMemos: number = 20
    UseCacheFirst: boolean = false
    UseCacheFileNum: number = 5
}

export const LocalSetting = defineStore('setting', () => {
    const frontSize = ref("14px");
    const mention: Ref<Array<{ label: string, value: string }>> = ref([]);
    const recentEditList: Ref<string[]> = ref([]);
    const allFileKeyList: Ref<string[]> = ref([]);
    const lastInput: Ref<string> = ref("");
    const delMemosList: Ref<string[]> = ref([]);
    let timer: string | number | NodeJS.Timeout | undefined = undefined;
    const localSetting: Ref<LocalSettingsClass> = ref(JSON.parse(localStorage.getItem('LocalSetting') || '{}'));
    recentEditList.value = JSON.parse(localStorage.getItem('recentEditList') || '["test.md"]');
    allFileKeyList.value = JSON.parse(localStorage.getItem('AllFileKeyList') || '["test.md"]');
    lastInput.value = localStorage.getItem('lastInput') || "";
    delMemosList.value = JSON.parse(localStorage.getItem('delMemosList') || '[]');

    onMounted(() => {
        getMention()
    })

    function getMention() {
        ObcsapiMentionGet().then(obj => {
            console.log("Load Mention")
            if (obj.tags != null) {
                mention.value = []
                obj.tags.forEach((val: string) => {
                    mention.value.push({ label: val, value: val });
                })
            }
        });
    }

    frontSize.value = JSON.parse(localStorage.getItem("theme") || "{}").frontSize

    function lastInputPush(text: string) {
        clearTimeout(timer);
        timer = setTimeout(function () {
            LocalSetting().lastInput = text;
            localStorage.setItem("lastInput", text);
        }, 1000);
    }

    function delMemosListPush(text: string) {
        delMemosList.value.unshift(text);
        localStorage.setItem('delMemosList', JSON.stringify(delMemosList.value));
    }


    return { mention, frontSize, getMention, recentEditList, 
        allFileKeyList, localSetting, 
        lastInput, lastInputPush,
        delMemosListPush,delMemosList
    }
})
