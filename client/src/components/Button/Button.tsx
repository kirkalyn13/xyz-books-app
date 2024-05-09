import React from "react";

interface ButtonProps {
    text: string;
    color?: string;
    hoverColor?: string;
    handleOnClick: Function;
    disable?: boolean;
}

const Button = ({text, 
    color = "bg-amber-500",
    hoverColor = "bg-amber-600",
    handleOnClick,
    disable = false,
    }: ButtonProps) => {
  return (
    <button 
        disabled={disable}
        className={`${!disable ? color : "bg-gray-300"} px-8 py-2 rounded-lg hover:${hoverColor} focus:outline-none`}
        onClick={() => handleOnClick()}>
        {text}
    </button>
  )
}

export default Button