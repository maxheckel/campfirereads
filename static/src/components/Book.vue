<template>
  <a :href="bookHref()" class="mt-8 mb-4 cover">
    <div class="flex items-center justify-center center">
      <img v-if="imageUrl() != ''" class="w-full transition-transform" :src="imageUrl()">
      <div v-else class="w-full h-[250px] bg-gray-200 block relative text-4xl overflow-hidden font-bold text-gray-300">{{props.book.volumeInfo.title}}</div>
    </div>
    <div class="w-full text-center px-2 mt-4">
      <span class="font-bold block"> {{props.book.volumeInfo.title}}</span>
      By: {{props.book.volumeInfo.authors.join(", ").trim(", ")}}
    </div>
  </a>

</template>

<script setup>
const props = defineProps({
  book: Object
})
console.log(props.book)
function imageUrl(){
  return props.book?.volumeInfo.imageLinks.thumbnail.replace("edge=curl", "")
}
function bookHref(){
  const isbn = props.book.volumeInfo.industryIdentifiers.find((ii) => ii.type == 'ISBN_13').identifier
  return '/book/'+isbn;
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