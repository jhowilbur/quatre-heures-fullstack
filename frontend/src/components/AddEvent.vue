<template>
  <div class="submit-form">
    <div v-if="!submitted">
      <div class="form-group">
        <label for="title">Title</label>
        <input
          type="text"
          class="form-control"
          id="title"
          required
          v-model="event.title"
          name="title"
        />
      </div>

      <div class="form-group">
        <label for="description">Description</label>
        <textarea
          class="form-control large-input"
          id="description"
          required
          v-model="event.description"
          name="description"
        />
      </div>

      <div class="form-group">
        <label>Start Date</label>
        <date-picker v-model="event.initial_date"></date-picker>
      </div>

      <div class="form-group">
        <label>End Date</label>
        <date-picker v-model="event.final_date"></date-picker>
      </div>

      <button @click="saveEvent" class="btn btn-success">Submit</button>
    </div>

    <div v-else>
      <h4>You submitted successfully!</h4>
      <button class="btn btn-success" @click="newEvent">Done</button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import EventDataService from "@/services/EventDataService";
import Event from "@/types/Event";
import ResponseData from "@/types/ResponseData";
import DatePicker from "vue3-datepicker";

export default defineComponent({
  name: "add-event",
  components: {
    DatePicker,
  },
  data() {
    return {
      event: {
        id: null,
        title: "",
        description: "",
        initial_date: null as Date | null,
        final_date: null as Date | null,
      } as Event,
      submitted: false,
    };
  },
  methods: {
    saveEvent() {
      if (
        this.event.final_date &&
        this.event.initial_date &&
        this.event.final_date < this.event.initial_date
      ) {
        alert("End date cannot be earlier than the start date.");
        return;
      }

      let data = {
        title: this.event.title,
        description: this.event.description,
        initial_date: this.event.initial_date,
        final_date: this.event.final_date,
      };

      EventDataService.create(data)
        .then((response: ResponseData) => {
          this.event.id = response.data.id;
          console.log(response.data);
          this.submitted = true;
        })
        .catch((e: Error) => {
          console.log(e);
        });
    },

    newEvent() {
      this.submitted = false;
      this.event = {} as Event;
    },
  },
});
</script>

<style>
.submit-form {
  max-width: 300px;
  margin: auto;
}
</style>
