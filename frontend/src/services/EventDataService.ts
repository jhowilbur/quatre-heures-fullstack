import http from "@/http-common";

/* eslint-disable */
class EventDataService {
  getAll(): Promise<any> {
    return http.get("/events");
  }

  get(id: any): Promise<any> {
    return http.get(`/events/${id}`);
  }

  create(data: any): Promise<any> {
    return http.post("/events", data);
  }
}

export default new EventDataService();
