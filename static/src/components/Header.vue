<template>
  <h2 class="flex items-center text-4xl uppercase font-bold text-np-dark-brown">
    <img class="w-6 inline-block mx-4" :src="props.iconPath"> -{{ props.text }}-<img class="w-6 inline-block mx-4"
                                                                                 :src="props.iconPath">
  </h2>
</template>

<script setup>
const props = defineProps({
  text: String,
  iconPath: String
})
</script>

<style scoped>

</style>