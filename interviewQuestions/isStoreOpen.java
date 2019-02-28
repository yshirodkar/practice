import java.time.DayOfWeek;
import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.List;

public class Solution {

    static class HoursEntry {
        int[] days;
        int startTime;
        int endTime;

        public HoursEntry(int[] days, int startTime, int endTime) {
            this.days = days;
            this.startTime = startTime;
            this.endTime = endTime;
        }
        //Getters omitted for brevity
    }

    public static void main(String[] args) {

        /*
          monday: 10am -> 4pm
          tuesday: 10am -> 4pm
          wednesday: 11am -> 8pm
          thursday: 11am -> 8pm
          friday: 11am -> 2am
          saturday: 5pm -> 4am
        */
        List<HoursEntry> exampleHours = Arrays.asList(
            new HoursEntry(new int[]{1, 2}, 10*60, 16*60),
            new HoursEntry(new int[]{3, 4}, 11*60, 20*60),
            new HoursEntry(new int[]{5}, 11*60, 24*60),
            new HoursEntry(new int[]{6}, 0*60, 2*60),
            new HoursEntry(new int[]{6}, 17*60, 24*60),
            new HoursEntry(new int[]{7}, 0*60, 4*60)
        );
        //Sunday, 2am
        LocalDateTime exampleTime = LocalDateTime.of(2019, 2, 16, 2, 0);

        System.out.println(isStoreOpen(exampleHours, exampleTime));
    }

    public static boolean isStoreOpen(List<HoursEntry> storeHours, LocalDateTime dt){
        int day = dayOfWeek(dt);
        int day_time = minutesSinceMidnight(dt);

        for (HoursEntry storeHour : storeHours) {
          if (storeHour.startTime <= day_time && day_time <= storeHour.endTime) {
                // Arrays.asList(yourArray).contains(yourValue)
                for (int i = 0; i < storeHour.days.length; i++) {
                  // System.out.printf("day: %d", storeHour.days[i]);
                  if (storeHour.days[i] == day) {
                      return true;
                  }
                }
            }
        }
        return false;
    }

    // Helper functions
    private static int minutesSinceMidnight(LocalDateTime dt){
        return dt.getMinute() + (dt.getHour() * 60);
    }

    private static int dayOfWeek(LocalDateTime dt){
        //1-indexed; Monday: 1 -> Sunday: 7
        return dt.getDayOfWeek().getValue();
    }

}