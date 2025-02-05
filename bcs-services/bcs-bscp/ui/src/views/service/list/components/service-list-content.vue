<script setup lang="ts">
  import { ref, computed, onMounted } from 'vue'
  import { storeToRefs } from 'pinia'
  import { useI18n } from "vue-i18n";
  import { Plus } from 'bkui-vue/lib/icon'
  import { useGlobalStore } from '../../../../store/global'
  import { useUserStore } from '../../../../store/user'
  import { getAppList, getAppsConfigData } from '../../../../api/index'
  import { IAppItem, IAppListQuery } from '../../../../../types/app'
  import Card from './card.vue'
  import CreateService from './create-service.vue'
  import EditService from './edit-service.vue'

  const { spaceId } = storeToRefs(useGlobalStore())
  const { userInfo } = storeToRefs(useUserStore())
  const { t } = useI18n();

  const props = defineProps<{
    type: string
  }>()

  const serviceList= ref<IAppItem[]>([])
  const isLoading = ref(true)
  const searchStr = ref('')
  const isCreateServiceOpen = ref(false)
  const isEditServiceOpen = ref(false)
  const editingService = ref<IAppItem>({
    id: 0,
    biz_id: 0,
    space_id: '',
    spec: {
      name: "",
      deploy_type: "",
      config_type: "",
      mode: "",
      memo: "",
      reload: {
        file_reload_spec: {
          reload_file_path: ""
        },
        reload_type: ""
      }
    },
    revision: {
      creator: "",
      reviser: "",
      create_at: "",
      update_at: "",
    }
  })
  const pagination = ref({
    current: 1,
    limit: 50,
    count: 0,
  });

  // 查询条件
  const filters = computed(() => {
    const { current, limit } = pagination.value;

    const rules: IAppListQuery = {
      start: (current - 1) * limit,
      limit: limit,
    };
    if (searchStr.value) {
      rules.name = searchStr.value
    }
    if (props.type === 'mine') {
      rules.operator = userInfo.value.username
    }
    return rules
  });
  const isEmpty = computed(() => {
    return serviceList.value.length === 0;
  });

  onMounted(() => {
    loadAppList()
  })

  // 加载服务列表
  const loadAppList = async () => {
    isLoading.value = true;
    try {
      const bizId = spaceId.value
      const resp = await getAppList(bizId, filters.value)
      if (resp.details.length > 0) {
        const appIds = resp.details.map((item: IAppItem) => item.id)
        const appsConfigData = await getAppsConfigData(bizId, appIds)
        resp.details.forEach((item: IAppItem, index: number) => {
          const { count, update_at } = appsConfigData.details[index]
          item.config = { count, update_at }
        })
      }
      // @ts-ignore
      serviceList.value = resp.details
      // @ts-ignore
      pagination.value.count = resp.count
    } catch (e) {
      console.error(e)
    } finally {
        isLoading.value = false;
    }
  };

  const handleNameInputChange = (val: string) => {
    if (!val) {
      handleSearch()
    }
  }

  // 编辑服务
  const handleEditService = (service: IAppItem) => {
    editingService.value = service
    isEditServiceOpen.value = true
  }

  // 搜索服务
  const handleSearch = () => {
    pagination.value.current = 1
    loadAppList()
  }

  // 删除服务后更新列表
  const handleDeletedUpdate = () => {
    if (serviceList.value.length === 1 && pagination.value.current > 1) {
      pagination.value.current -= 1
    }
    loadAppList();
  }

  const handleLimitChange = (limit: number) => {
    pagination.value.limit = limit
    loadAppList();
  }
</script>
<template>
    <div class="service-list-content">
      <div class="head-section">
        <bk-button theme="primary" @click="isCreateServiceOpen = true">
          <Plus class="create-icon" />
          {{ t("新建服务") }}
        </bk-button>
        <div class="head-right">
          <bk-input
            class="search-app-name"
            type="search"
            v-model="searchStr"
            :placeholder="t('服务名称')"
            :clearable="true"
            @change="handleNameInputChange"
            @enter="handleSearch"
            @clear="handleSearch">
          </bk-input>
        </div>
      </div>
      <div class="content-body">
        <bk-loading style="height: 100%;" :loading="isLoading">
          <template v-if="!isLoading && isEmpty">
            <bk-exception
              class="exception-wrap-item"
              type="empty"
              :description="t('你尚未创建或加入任何服务')">
              <div class="exception-actions">
                <bk-button text theme="primary" @click="isCreateServiceOpen = true">
                  {{ t("立即创建") }}
                </bk-button>
                <span class="divider-middle"></span>
                <!-- <bk-button text theme="primary">{{ t("申请权限") }}</bk-button> -->
              </div>
            </bk-exception>
          </template>
          <template v-else>
            <div class="serving-list">
              <Card
                v-for="service in serviceList"
                class="service-item"
                :key="service.id"
                :service="service"
                @edit="handleEditService"
                @update="handleDeletedUpdate" />
            </div>
            <bk-pagination
              v-model="pagination.current"
              class="service-list-pagination"
              location="left"
              :layout="['total', 'limit', 'list']"
              :count="pagination.count"
              :limit="pagination.limit"
              @change="loadAppList"
              @limit-change="handleLimitChange" />
          </template>
          </bk-loading>
      </div>
      <CreateService v-model:show="isCreateServiceOpen" @reload="loadAppList" />
      <EditService v-model:show="isEditServiceOpen" :service="editingService" />
    </div>
</template>
<style lang="scss" scoped>
  .service-list-content {
    height: 100%;
  }
  .head-section {
    display: flex;
    justify-content: space-between;
    margin: 0 auto;
    padding: 16px 0;
    width: 1200px;
    .create-icon {
      font-size: 22px;
    }
    .head-right {
      display: flex;
      .search-app-name {
        margin-left: 16px;
        width: 240px;
      }
    }
  }
  .content-body {
    margin: 0 auto;
    padding-bottom: 24px;
    width: 1216px;
    height: calc(100% - 64px);
    overflow: auto;
    .serving-list {
      display: flex;
      flex-wrap: wrap;
      align-content: flex-start;
      :deep(.bk-exception-description) {
        margin-top: 5px;
        font-size: 12px;
        color: #979ba5;
      }
      :deep(.bk-exception-footer) {
        margin-top: 5px;
      }
      .exception-actions {
        display: flex;
        font-size: 12px;
        color: #3a84ff;
        .divider-middle {
          display: inline-block;
          margin: 0 16px;
          width: 1px;
          height: 16px;
          background: #dcdee5;
        }
      }
    }
  }
  .service-list-pagination {
    padding: 0 8px;
    :deep(.bk-pagination-list.is-last) {
      margin-left: auto;
    }
  }
</style>