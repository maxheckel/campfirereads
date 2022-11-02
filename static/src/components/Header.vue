<template>
  <h2 class="flex items-center text-2xl md:text-4xl uppercase font-bold text-np-dark-brown">
    <img :class="getImageClasses()" class="hidden md:block inline-block mx-4" :src="props.iconPath"> -{{ props.text }}-<img :class="getImageClasses()" class="hidden md:block  inline-block mx-4"
                                                                                 :src="props.iconPath">
  </h2>
</template>

<script setup>
const props = defineProps({
  text: String,
  iconPath: String,
  widthOverride: String,
  invert: Boolean
})

const getImageClasses = () => {
  let resp = {}
  resp['w-4'] = true
  if(props.widthOverride){
    resp['w-4'] = false;
    resp[props.widthOverride] = true;
  }
  if(props.invert){
    resp['invert'] = true;
  }
  return resp;
}
</script>

<style scoped>

</style>