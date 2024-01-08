<template>
  <a-tree v-model:selectedKeys="selectedKeys" :tree-data="treeData" show-icon default-expand-all>
    <template #switcherIcon="{ switcherCls }"><down-outlined :class="switcherCls" /></template>
    <template #icon="{ key, selected }">
      <template v-if="key === '0-0'">
        <smile-outlined />
      </template>
      <template v-else-if="key === '0-0-0'">
        <smile-outlined />
      </template>
      <template v-else>
        <frown-filled v-if="selected" />
        <frown-outlined v-else />
      </template>
    </template>
  </a-tree>
</template>
<script setup>
import { ref } from 'vue';
const treeData = [
  {
    title: 'parent 1',
    key: '0-0',
    children: [
      {
        title: 'leaf',
        key: '0-0-0',
      },
      {
        title: 'leaf',
        key: '0-0-1',
      },
    ],
  },
];
const selectedKeys = ref(['0-0-0']);
</script>
