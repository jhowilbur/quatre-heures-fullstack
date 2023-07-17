import { mount } from '@vue/test-utils';
import EventComponent from '@/components/EventComponent.vue';

// TODO - Finish this test

describe('EventComponent', () => {
  it('renders the event details when currentEvent is populated', async () => {
    const wrapper = mount(EventComponent, {
      props: {
        currentEvent: {
          id: 1,
          title: 'Sample Event',
          description: 'This is a sample event',
          initial_date: '2023-07-01',
          final_date: '2023-07-10',
        },
      },
    });

    await wrapper.vm.$nextTick();

    expect(wrapper.find('.edit-form').exists()).toBe(true);
    expect(wrapper.find('#title').element.value).toBe('Sample Event');
    expect(wrapper.find('#description').element.value).toBe('This is a sample event');
    expect(wrapper.find('.form-group').text()).toContain('Initial Date: 2023-07-01');
    expect(wrapper.find('.form-group').text()).toContain('Final Date: 2023-07-10');
    expect(wrapper.find('p').text()).toBe('');
  });

  it('renders the "Please click on an Event..." message when currentEvent is not populated', () => {
    const wrapper = mount(EventComponent, {
      props: {
        currentEvent: {},
      },
    });

    expect(wrapper.find('.edit-form').exists()).toBe(false);
    expect(wrapper.find('p').text()).toBe('Please click on a Event...');
  });
});
