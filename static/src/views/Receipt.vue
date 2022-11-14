<template>

</template>

<script setup>
import {onMounted, reactive} from "vue";
  import {useRoute} from "vue-router/dist/vue-router";

  const route = useRoute();
  const id = route.params.id

  const data = reactive({
    receipt: {},
    loading: true
  })

  onMounted(() => {

    fetch(import.meta.env.VITE_API_HOST + "receipt/" + id)
        .then((response) => response.json())
        .then((resp) => {
          data.loading = false;
          data.receipt = resp
        });
  })

</script>

<style scoped>

</style>