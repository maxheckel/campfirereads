<template>
  <div class="w-full text-center">
    <img :src="image" class="rounded-full block relative mx-auto mb-4 w-1/2 h-1/2">
    <div class="text-xl mb-2 font-bold">{{name}}</div>
    <span class="italic mb-2 block">{{title}}</span>
    <span class="block font-normal">{{description}}</span>
  </div>
</template>

<script setup>
  const props = defineProps({
    name: String,
    image: String,
    title: String,
    description: String
  })
</script>

<style scoped>

</style>