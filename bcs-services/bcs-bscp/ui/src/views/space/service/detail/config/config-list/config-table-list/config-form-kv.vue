<template>
  <bk-form ref="formRef" form-type="vertical" :model="localVal" :rules="rules">
    <bk-form-item label="配置项名称" property="key" :required="true">
      <bk-input v-model="localVal.key" :disabled="editable || view" @change="change" />
    </bk-form-item>
    <bk-form-item label="数据类型" property="kv_type" :required="true" :description="typeDescription">
      <bk-radio-group v-model="localVal.kv_type">
        <bk-radio
          v-for="kvType in CONFIG_KV_TYPE"
          :key="kvType.id"
          :label="kvType.id"
          :disabled="appData.spec.data_type !== 'any' || editable || view"
          >{{ kvType.name }}</bk-radio
        >
      </bk-radio-group>
    </bk-form-item>
    <bk-form-item label="配置项值" property="value" :required="true">
      <bk-input
        v-if="localVal.kv_type === 'string' || localVal.kv_type === 'number'"
        v-model.trim="localVal!.value"
        @change="change"
        :disabled="view"
      />
      <KvConfigContentEditor
        v-else
        :languages="localVal.kv_type"
        :content="localVal.value"
        :editable="!view"
        @change="handleStringContentChange"
      />
    </bk-form-item>
  </bk-form>
</template>

<script lang="ts" setup>
import { ref, onMounted, computed } from 'vue';
import { CONFIG_KV_TYPE } from '../../../../../../../constants/config';
import KvConfigContentEditor from '../../components/kv-config-content-editor.vue';
import { IConfigKvEditParams } from '../../../../../../../../types/config';
import useServiceStore from '../../../../../../../store/service';
import { storeToRefs } from 'pinia';
import { validateJSON, validateXML, validateYAML } from '../../../../../../../utils/kv-validate';

const serviceStore = useServiceStore();
const { appData } = storeToRefs(serviceStore);

const props = withDefaults(
  defineProps<{
    config: IConfigKvEditParams;
    editable?: boolean;
    view?: boolean;
    bkBizId: string;
    id: number; // 服务ID或者模板空间ID
    isTpl?: boolean; // 是否未模板配置文件，非模板配置文件和模板配置文件的上传、下载接口参数有差异
  }>(),
  {
    editable: false,
    view: false,
  },
);

const formRef = ref();
const localVal = ref({
  ...props.config,
});

const typeDescription = computed(() => {
  if (appData.value.spec.data_type !== 'any' && !props.editable && !props.view) {
    return `已限制该服务下所有配置项数据类型为${appData.value.spec.data_type}，如需其他数据类型，请调整服务属性下的数据类型`;
  }
  return '';
});

const rules = {
  value: [
    {
      validator: (value: string) => {
        if (localVal.value.kv_type === 'number') {
          return /^-?\d+(\.\d+)?$/.test(value);
        }
        return true;
      },
      message: '配置项值不为数字',
    },
  ],
};

// 新建文件任意类型默认选中string
onMounted(() => {
  if (!props.editable && !props.view) {
    localVal.value.kv_type = appData.value.spec.data_type! === 'any' ? 'string' : appData.value.spec.data_type!;
  }
});

const validate = async () => {
  await formRef.value.validate();
  switch (localVal.value.kv_type) {
    case 'json':
      return validateJSON(localVal.value.value);
    case 'xml':
      return validateXML(localVal.value.value);
    case 'yaml':
      return validateYAML(localVal.value.value);
  }
  return true;
};


const emits = defineEmits(['change']);


const handleStringContentChange = (val: string) => {
  localVal.value!.value = val;
  change();
};

const change = () => {
  emits('change', localVal.value);
};

defineExpose({ validate });
</script>

<style scoped lang="scss"></style>
