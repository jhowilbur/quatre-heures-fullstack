<template>
  <div v-if="currentEvent.id" class="edit-form">
    <h4>Event</h4>
    <form>
      <div class="form-group">
        <label for="title">Title</label>
        <input
          type="text"
          class="form-control"
          id="title"
          v-model="currentEvent.title"
        />
      </div>
      <div class="form-group">
        <label for="description">Description</label>
        <textarea
          type="text"
          class="form-control"
          id="description"
          v-model="currentEvent.description"
        />
      </div>

      <div class="form-group">
        <label><strong>Initial Date:</strong></label>
        {{ currentEvent.initial_date }}
        <label><strong>Final Date:</strong></label>
        {{ currentEvent.final_date }}
      </div>
    </form>

    <p>{{ message }}</p>
  </div>

  <div v-else>
    <br />
    <p>Please click on a Event...</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import EventDataService from "@/services/EventDataService";
import Event from "@/types/Event";
import ResponseData from "@/types/ResponseData";

export default defineComponent({
  name: "event",
  data() {
    return {
      currentEvent: {} as Event,
      message: "",
    };
  },
  methods: {
    getEvent(id: any) {
      EventDataService.get(id)
        .then((response: ResponseData) => {
          this.currentEvent = response.data;
          console.log(response.data);
        })
        .catch((e: Error) => {
          console.log(e);
        });
    },
  },
  mounted() {
    this.message = "";
    this.getEvent(this.$route.params.id);
  },
});
</script>

<style>
.edit-form {
  max-width: 300px;
  margin: auto;
}
</style>
