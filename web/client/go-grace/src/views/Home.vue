<template>
	<div class="cascader">
		<a-cascader
			v-model:value="state.value"
			multiple
			:options="options"
			placeholder="Please select"
			suffix-icon="Shopping Around"
      style="width: 300px;"
		>
			<template #tagRender="data">
				<a-tag :key="data.value" color="blue">{{ data.label }}</a-tag>
			</template>
		</a-cascader>
	</div>
</template>
<script setup>
import { ref, reactive } from "vue"
const options = [
	{
		value: "zhejiang",
		label: "Zhejiang",
		children: [
			{
				value: "freezeKey",
				label: "freeze"
			},
			{
				value: "fruitsKey",
				label: "fruits"
			}
		]
	},
	{
		value: "Chinese delicious food",
		label: "中国美食",
		children: [
			{
				value: "key3",
				label: "月饼",
				children: [
					{
						value: "key4",
						label: "蛋黄馅"
					},
					{
						value: "key5",
						label: "五仁馅"
					}
				]
			}
		]
	}
]
const state = reactive({
  value: []
})
</script>

<style>
.cascader {
	width: 100%;
}
</style>
