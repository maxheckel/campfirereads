<template>
  <div class="p-2 font-bold cursor-pointer px-8 inline-block rounded rounded-md text-lg bg-np-dark-brown text-np-yellow">{{text}}</div>
</template>

<script setup>
  const props = defineProps({
    text: String
  })
</script>

<style scoped>

</style>