pub mod rest;
pub mod websocket;

pub use rest::client::{DnseClient, Config};
// We will export more things as we build them
