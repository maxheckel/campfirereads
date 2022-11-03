<template>
  <div class="flex my-4 items-center">
    <b class="float-left">{{ label }}</b>
    <div class="dot"></div>
    <div class="ml-auto float-left">
      {{value}}
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  label: String,
  value: String
})
</script>

<style scoped>

.dot{
  border-bottom: dotted 1px orange;
  width: 100%;
  float: left;
  position: relative;
  display: block;
  height: 12px;
  margin: 0 5px 0 5px;
}
</style>