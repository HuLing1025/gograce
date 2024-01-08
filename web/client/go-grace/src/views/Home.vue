  <template>
	<a-tree
	  v-model:expandedKeys="expandedKeys"
	  v-model:selectedKeys="selectedKeys"
	  v-model:checkedKeys="checkedKeys"
	  checkable
	  :tree-data="treeData"
	>
	  <template #title="{ title, key }">
		<span v-if="key === '0-0-1-0'" style="color: #1890ff">{{ title }}</span>
		<template v-else>{{ title }}</template>
	  </template>
	</a-tree>
  </template>
  <script setup>
  import { ref, watch } from 'vue';
  import {get, post} from "../utils/axios";


 	const treeData = ref([]);
	get("http://localhost:8080/v1/tree", "").then((response) => {
		console.log(response.data);
		console.log(response.data.length);

		for (var i = 0; i < response.data.length; i++){
			var directory = ref([]);
			directory.title = response.data[i].directory;
			directory.key = response.data[i].directory;

			for(var j = 0; j < response.data[i].files.length; j++){
				var file = ref([]);
				file.title = response.data[i].files[j].file;
				file.key = response.data[i].files[j].file;
				for(var k = 0; k < response.data[i].files[j].functions.length; k++){
					file.functions.push(response.data[i].files[j].functions[k].name)
				}
				for(var k = 0; k < response.data[i].files[j].methods.length; k++){
					file.methods.push(response.data[i].files[j].methods[k].struct + response.data[i].files[j].methods[k].name)
				}
				directory.children.push(file)
			}

			treeData.push(directory)
		}
	});
	
//   const treeData = [
// 	{
// 	  title: 'parent 1',
// 	  key: '0-0',
// 	  children: [
// 		{
// 		  title: 'parent 1-0',
// 		  key: '0-0-0',
// 		  disabled: true,
// 		  children: [
// 			{
// 			  title: 'leaf',
// 			  key: '0-0-0-0',
// 			  disableCheckbox: true,
// 			},
// 			{
// 			  title: 'leaf',
// 			  key: '0-0-0-1',
// 			},
// 		  ],
// 		},
// 		{
// 		  title: 'parent 1-1',
// 		  key: '0-0-1',
// 		  children: [
// 			{
// 			  key: '0-0-1-0',
// 			  title: 'sss',
// 			},
// 		  ],
// 		},
// 	  ],
// 	},
//   ];
  const expandedKeys = ref(['0-0-0', '0-0-1']);
  const selectedKeys = ref(['0-0-0', '0-0-1']);
  const checkedKeys = ref(['0-0-0', '0-0-1']);
  watch(expandedKeys, () => {
	console.log('expandedKeys', expandedKeys);
  });
  watch(selectedKeys, () => {
	console.log('selectedKeys', selectedKeys);
  });
  watch(checkedKeys, () => {
	console.log('checkedKeys', checkedKeys);
  });
  </script>