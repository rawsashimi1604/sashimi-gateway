// Generate time data for chart js
export function generateTimeData(
  intervalInMinutes: number,
  length: number
): Date[] {
  return Array.from({ length })
    .map((_, i) => {
      const date = new Date();
      date.setMinutes(date.getMinutes() - i * intervalInMinutes);
      return date;
    })
    .reverse();
}
