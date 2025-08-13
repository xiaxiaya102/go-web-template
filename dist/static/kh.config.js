import apiUrl from "./apiUrl";
export default {
  host: {
    rest: "",
    basePath: "",
    gisType: "天地图", //高德,//天地图
    publicKeyStr:
      "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmzJfX5YZIZMj5Qkgq+ph3apVDXE2DtYBtw4GkjzQjhGCDec18Ddx/fHgFCupRNWqsaGPwhqkFV65k/M6F0lRgjfyis07tAgRZVASI6l/89jNGWFXAik7jo63rxqZFRMr1GYu7H9SjCPT5QogJ8ARjdh1hKcfV0KBBDhyGOSZE4Gkoour1ZEHtffiXSucQ/n/EArpXTWEjhjXQ6sX0gN560ua4KJu3KXIOMAaFPyRgy7wxFXQOEgMySBfdvsnVjeMIKXIrGkqhzTxEr+RO8e98qldPxs/bjjIZu12QYUddGxuv292XnD1u5MJIDYSrAZmk++ApXOMYT/IPFtmXfkWCwIDAQAB",
  },
  api: {
    ...apiUrl,
  },
  msg: {
    no_login: {
      title: "登录失效",
      content: "请重新登录",
    },
  },
  components: {
    ListFormCustom: {
      props: ["listBusinessId"],
    },
    ViewFormCustom: {
      props: ["businessId"],
    },
    FormGroupAdd: {
      props: ["visible"],
    },
    ThemeListShow: {
      props: ["currentType"],
    },
  },
};
