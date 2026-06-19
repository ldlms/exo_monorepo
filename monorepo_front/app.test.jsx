import { render} from '@testing-library/react'
import { describe, it, expect } from 'vitest'
import App from './src/App'

describe('App', () => {
  it('s\'affiche sans planter', () => {
    render(<App />)
    expect(document.body).toBeInTheDocument()
  })
})