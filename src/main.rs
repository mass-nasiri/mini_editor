// src/main.rs

use tao::{
    event::{Event, StartCause, WindowEvent},
    event_loop::{ControlFlow, EventLoop},
    window::WindowBuilder,
};
use wry::WebViewBuilder;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Include the HTML and editor assets directly inside the binary for speed and portability
    let html_content = include_str!("index.html");

    let event_loop = EventLoop::new();
    let window = WindowBuilder::new()
        .with_title("Mini HTML/CSS/JS/TS Live Editor")
        .with_inner_size(tao::dpi::LogicalSize::new(1200.0, 800.0))
        .build(&event_loop)?;

    // Safely wrap the window inside the WebViewBuilder setup
    let _webview = WebViewBuilder::new(&window)
        .with_html(html_content)?
        .build()?;

    event_loop.run(move |event, _, control_flow| {
        *control_flow = ControlFlow::Wait;

        match event {
            Event::NewEvents(StartCause::Init) => println!("Editor successfully started!"),
            Event::WindowEvent {
                event: WindowEvent::CloseRequested,
                ..
            } => *control_flow = ControlFlow::Exit,
            _ => (),
        }
    });
}
