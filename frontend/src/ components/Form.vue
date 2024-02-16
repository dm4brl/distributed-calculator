<template>
  <div>
    <input type="text" v-model="expression" placeholder="Enter expression">
    <button @click="submitExpression">Submit</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      expression: ''
    };
  },
  methods: {
    submitExpression() {
      // Implement logic to submit expression to the backend
    }
  }
};
</script>
