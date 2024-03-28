<script setup lang="tsx">
import { Form } from '@/components/Form'
import { reactive, ref, onMounted, computed } from 'vue'
import { useValidator } from '@/hooks/web/useValidator'
import { useI18n } from '@/hooks/web/useI18n'
import { useForm } from '@/hooks/web/useForm'
import { ContentWrap } from '@/components/ContentWrap'
import { useAppStore } from '@/store/modules/app'
import { SelectOption, RadioOption, CheckboxOption, FormSchema } from '@/components/Form'
import {
  ElOption,
  ElOptionGroup,
  ElRadio,
  ElRadioButton,
  ElCheckbox,
  ElCheckboxButton,
  ElInput,
  ElMessage,
  ElMessageBox,
  ElIcon,
  ElSteps,
  ElStep,
  ElButton
} from 'element-plus'
import { getDictOneApi } from '@/api/common'
import { Icon } from '@/components/Icon'
import { BaseButton } from '@/components/Button'
import { useIcon } from '@/hooks/web/useIcon'
import axios from 'axios'
import { propTypes } from '@/utils/propTypes'
import { personalCertApi, personalCertStatusApi } from '@/api/cert'
import { PersonalCert } from '@/api/cert/types'
import { log } from 'console'
import { useRouter } from 'vue-router'
import { createImageViewer } from '@/components/ImageViewer'

const emit = defineEmits(['search', 'reset', 'expand'])
const { formMethods, formRegister } = useForm()
const { getElFormExpose } = formMethods
const appStore = useAppStore()
const { required } = useValidator()
const { t } = useI18n()
defineProps({
  showSearch: propTypes.bool.def(true),
  showReset: propTypes.bool.def(true),
  showExpand: propTypes.bool.def(false),
  visible: propTypes.bool.def(true),
  searchLoading: propTypes.bool.def(false),
  resetLoading: propTypes.bool.def(false)
})
const isMobile = computed(() => appStore.getMobile)

const router = useRouter()
const checkCertStatus = async () => {
  try {
    const resp = await personalCertStatusApi()
    const status = resp.status
    if (status === 0) {
      console.log()
    } else if (status === 2) {
      router.push({ path: '/certification/cwaiting' })
    } else if (status === 1) {
      router.push({ path: '/certification/ccompleted' })
    }
  } catch (e) {}
}

const restaurants = ref<Recordable[]>([])
const querySearch = (queryString: string, cb: Fn) => {
  const results = queryString
    ? restaurants.value.filter(createFilter(queryString))
    : restaurants.value
  // call callback function to return suggestions
  cb(results)
}
let timeout: NodeJS.Timeout
const querySearchAsync = (queryString: string, cb: (arg: any) => void) => {
  const results = queryString
    ? restaurants.value.filter(createFilter(queryString))
    : restaurants.value

  clearTimeout(timeout)
  timeout = setTimeout(() => {
    cb(results)
  }, 3000 * Math.random())
}
const createFilter = (queryString: string) => {
  return (restaurant: Recordable) => {
    return restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
  }
}
const loadAll = () => {
  return [
    { value: 'vue', link: 'https://github.com/vuejs/vue' },
    { value: 'element', link: 'https://github.com/ElemeFE/element' },
    { value: 'cooking', link: 'https://github.com/ElemeFE/cooking' },
    { value: 'mint-ui', link: 'https://github.com/ElemeFE/mint-ui' },
    { value: 'vuex', link: 'https://github.com/vuejs/vuex' },
    { value: 'vue-router', link: 'https://github.com/vuejs/vue-router' },
    { value: 'babel', link: 'https://github.com/babel/babel' }
  ]
}
const handleSelect = (item: Recordable) => {
  console.log(item)
}
onMounted(() => {
  checkCertStatus()
  restaurants.value = loadAll()
})

const loading = ref(false)
const { getFormData, setValues } = formMethods

const onSuccessFront = async (response) => {
  console.log(response)
  if (response && response.code === 0 && response.card_path) {
    const formData = await getFormData<PersonalCert>()
    formData.card_front = response.card_path
  }
}

const onSuccessBack = async (response) => {
  if (response && response.code === 0 && response.card_path) {
    const formData = await getFormData<PersonalCert>()
    formData.card_back = response.card_path
  }
}

const onPreviewFront = async (response) => {
  if (response && response.response.code === 0 && response.response.card_path) {
    imageUrl.value = response.response.card_path
  }
}

const onPreviewBack = async (response) => {
  if (response && response.response.code === 0 && response.response.card_path) {
    imageUrl2.value = response.response.card_path
  }
}

const onSubmit = async () => {
  const formRef = await getElFormExpose()
  formRef?.validate(async (valid) => {
    if (valid) {
      loading.value = false
      const formData = await getFormData<PersonalCert>()
      const { field4, field5, ...rest } = formData
      const tem = {
        ...rest,
        card_front: field4[0].response.card_path,
        card_back: field5[0].response.card_path
      }
      try {
        loading.value = true
        personalCertApi(tem)
      } finally {
        loading.value = false
        router.push({ path: '/certification/cwaiting' })
      }
    }
  })
}

let id = 0

const imageUrl = ref('')
const imageUrl2 = ref('')

const schema = reactive<FormSchema[]>([])

const redirectToPersonUrl = () => {
  router.push('/certification/person')
}

const redirectToCorpUrl = () => {
  router.push('/certification/corp')
}
</script>

<template>
  <ContentWrap :title="t('certification.cert')">
    <!-- <el-steps style="max-width: auto" :active="1" align-center>
      <el-step title="填写认证材料" />
      <el-step title="提交认证审核" />
      <el-step title="实名认证完成" />
    </el-steps> -->
    <ElButton class="avatar-uploader-icon" @click="redirectToPersonUrl"> 个人认证</ElButton>
    <ElButton class="avatar-uploader-icon" @click="redirectToCorpUrl"> 企业认证</ElButton>
    <Form
      :schema="schema"
      @register="formRegister"
      label-width="auto"
      :label-position="isMobile ? 'top' : 'right'"
    />
  </ContentWrap>
</template>

<style lang="less">
.avatar {
  width: 200px;
}
.cell {
  height: 30px;
  padding: 3px 0;
  box-sizing: border-box;

  .text {
    position: absolute;
    left: 50%;
    display: block;
    width: 24px;
    height: 24px;
    margin: 0 auto;
    line-height: 24px;
    border-radius: 50%;
    transform: translateX(-50%);
  }

  &.current {
    .text {
      color: #fff;
      background: #626aef;
    }
  }

  .holiday {
    position: absolute;
    bottom: 0;
    left: 50%;
    width: 6px;
    height: 6px;
    background: var(--el-color-danger);
    border-radius: 50%;
    transform: translateX(-50%);
  }
}

.transfer-footer {
  padding: 6px 5px;
  margin-left: 15px;
}

.el-upload {
  position: relative;
  overflow: hidden;
  cursor: pointer;
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  transition: var(--el-transition-duration-fast);
}

.el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  width: 178px;
  height: 178px;
  font-size: 28px;
  color: #8c939d;
  text-align: center;
}
</style>
