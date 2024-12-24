import { onMounted, onUnmounted } from 'vue'

export default function useInterval(callback: () => void, time: number): void {
  let timer: NodeJS.Timeout | undefined = undefined

  onMounted(() => {
    timer = setInterval(callback, time)
  })

  onUnmounted(() => {
    clearInterval(timer)
  })
}
