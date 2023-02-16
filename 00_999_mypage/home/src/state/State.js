
// 辞書
export const WordDataInit = {
    code: "000",
    w: {
        ID: "",
        CreatedAt: "",
        UpdatedAt: "",
        DeletedAt: "",
        Word: "",
        Syllable: "",
        PhoneticSymbols: "",
        Info: "",
        InfosJp: [],
        InfosEn: "",
        VoiceLink: "",
        Memo: "",
    }
}

// 履歴データ
export const HistoryDataInit = {
    code: "000",
    h: [
        {
            ID: "",
            CreatedAt: "",
            UpdatedAt: "",
            DeletedAt: "",
            Word: "",
            Syllable: "",
            PhoneticSymbols: "",
            Info: "",
            InfosJp: "",
            InfosEn: "",
            VoiceLink: "",
            Memo: "",
        }
    ],
}

export const StsDataInit = {
    code: "000",
    dc: [
        {
            day: "",
            dayCount: 0
        },
    ],
    dm: [
        {
            month: "",
            monthCount: 0
        },
    ],
}


export const reducerWordFunc = (WordState, action) => {
    switch (action.type) {
        case 'success':
            return {
                ...WordState, ...action.value
            }
        default:
            return WordState
    }

}