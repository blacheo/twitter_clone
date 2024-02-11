<script setup>
defineProps({
    user: {
        type: String,
        required: true
    },
    content: {
    type: String,
    required: true
  }
})
</script>

<template>
    <h3>{{ user }}</h3>
    <p>{{ content }}</p>
</template>
