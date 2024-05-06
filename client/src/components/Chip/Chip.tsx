import React from 'react'

interface ChipProps {
    id: string
    name: string
    removeChip: Function
}
const Chip: React.FC<ChipProps> = ({ id, name, removeChip }) => {
  return (
    <div
        key={id}
        className="bg-slate-800 text-white py-1 px-2 rounded-lg flex items-center"
    >
    {name}
    <button
        className="ms-2 text-sm font-bold"
        onClick={() => removeChip(id)}
    >
        X
    </button>
    </div>
  )
}

export default Chip