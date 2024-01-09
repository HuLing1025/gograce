<template>
	<a-tree
		v-model:expandedKeys="expandedKeys"
		v-model:selectedKeys="selectedKeys"
		v-model:checkedKeys="checkedKeys"
		:tree-data="treeData"
	>
		<template #title="{ title, key }">
			<span v-if="key === '0-0-1-0'" style="color: #1890ff">{{ title }}</span>
			<template v-else>{{ title }}</template>
		</template>
	</a-tree>
</template>
<script setup>
import { ref, watch } from "vue"
import { get, post } from "../utils/axios"

const treeData = ref([])

const initOptions = (list) => {
	return list.map((item) => {
		let children = undefined

		// handel file.
		if (item.files && item.files.length) {
			children = initOptions(item.files)
			return {
				title: item.directory,
				key: item.directory,
				children,
			}
		}

		if (item.functions && item.functions.length) {
			children = initOptions(item.functions)
			if (!item.methods || !item.methods.length){
				return {
					title: item.name,
					key: item.name,
					children,
				}
			}
		}

		if (item.methods && item.methods.length) {
			var m = initOptions(item.methods)
			m.forEach(element => {
				children.push(element) 
			});
			return {
				title: item.name,
				key: item.name,
				children,
			}
		}

		// methods...
		if (item.struct){
			return {
				title: "[Method] " + item.struct + "." + item.name,
				key: item.struct + "." + item.name,
				children: null,
			}
		}

		// functions...
		if (item){
			return {
				title: "[Function] " + item,
				key: item,
				children: null,
			}
		}
		
	})
}

get("/v1/tree").then((response) => {
	if (response.code === 200) {
		console.log(response.data)
		treeData.value = initOptions(response.data)
	}
	console.log(treeData.value)
})

const expandedKeys = ref(["0-0-0", "0-0-1"])
const selectedKeys = ref(["0-0-0", "0-0-1"])
const checkedKeys = ref(["0-0-0", "0-0-1"])
watch(expandedKeys, () => {
	// console.log("expandedKeys", expandedKeys)
})
watch(selectedKeys, () => {
	// console.log("selectedKeys", selectedKeys)
})
watch(checkedKeys, () => {
	// console.log("checkedKeys", checkedKeys)
})
</script>
