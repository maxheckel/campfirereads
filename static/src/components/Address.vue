<template>
  <b>{{type}}</b>
  <div>{{address.street1}}</div>
  <div>{{address.street2}}</div>
  <div>{{address.city}}</div>
  <div>{{address.state}}</div>
  <div>{{address.zip}}</div>
</template>

<script setup>
  const props = defineProps({
    address: Object,
    type: String
  })
</script>

<style scoped>

</style>