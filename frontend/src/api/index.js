var socket=new WebSocket("ws://localhost:8080/ws")

let connect =()=>{
    let salidaTxt=document.getElementById("output")
    console.log("intentando conexion");

    socket.onopen=()=>{
        console.log("conectado con exito");
    }

    socket.onmessage=(msg)=>{
        salidaTxt.innerHTML+=msg.data
        console.log(msg.data);
    }
    socket.onclose=(event)=>{
        console.log("conexion de socket cerrada ",event);
    }
}

let senMsg=(msg)=>{
    console.log("enviendo mensaje",msg);
    socket.send(msg)
}

export {
    connect,senMsg
}