<template>
  <div class="mt-8 mb-4 cover">
    <div class="flex items-center justify-center center">
      <img class="w-full transition-transform" :src="imageUrl()">
    </div>
    <div class="w-full text-center px-2 mt-4">
      <span class="font-bold block"> {{props.book.volumeInfo.title}}</span>
      By: {{props.book.volumeInfo.authors.join(", ").trim(", ")}}
    </div>
  </div>

</template>

<script setup>
const props = defineProps({
  book: Object
})
function imageUrl(){
  return props.book.volumeInfo.imageLinks.thumbnail.replace("edge=curl", "")
}
</script>

<style scoped lang="scss">
  .cover{
    position: relative;
    &:hover{
        cursor:pointer;
        img{
          transform: scale(1.1);

        }
    }
  }

</style>