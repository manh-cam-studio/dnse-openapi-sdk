use base64::{engine::general_purpose::STANDARD as BASE64, Engine};
use hmac::{Hmac, Mac};
use sha1::Sha1;
use sha2::{Sha256, Sha384, Sha512};
use url::form_urlencoded;

pub fn build_signature(
    secret: &str,
    method: &str,
    path: &str,
    date_value: &str,
    algorithm: &str,
    nonce: Option<&str>,
    header_name: &str,
) -> Result<(String, String), String> {
    let header_key = header_name.to_lowercase();
    let headers_list = format!("(request-target) {}", header_key);

    let mut signature_string = format!(
        "(request-target): {} {}\n{}: {}",
        method.to_lowercase(),
        path,
        header_key,
        date_value
    );

    if let Some(n) = nonce {
        signature_string.push_str(&format!("\nnonce: {}", n));
    }

    let encoded_mac = match algorithm {
        "hmac-sha256" => {
            let mut mac = Hmac::<Sha256>::new_from_slice(secret.as_bytes()).map_err(|e| e.to_string())?;
            mac.update(signature_string.as_bytes());
            Result::<String, String>::Ok(BASE64.encode(mac.finalize().into_bytes()))
        },
        "hmac-sha384" => {
            let mut mac = Hmac::<Sha384>::new_from_slice(secret.as_bytes()).map_err(|e| e.to_string())?;
            mac.update(signature_string.as_bytes());
            Result::<String, String>::Ok(BASE64.encode(mac.finalize().into_bytes()))
        },
        "hmac-sha512" => {
            let mut mac = Hmac::<Sha512>::new_from_slice(secret.as_bytes()).map_err(|e| e.to_string())?;
            mac.update(signature_string.as_bytes());
            Result::<String, String>::Ok(BASE64.encode(mac.finalize().into_bytes()))
        },
        _ => {
            let mut mac = Hmac::<Sha1>::new_from_slice(secret.as_bytes()).map_err(|e| e.to_string())?;
            mac.update(signature_string.as_bytes());
            Result::<String, String>::Ok(BASE64.encode(mac.finalize().into_bytes()))
        }, // Fallback
    }?;

    let escaped = form_urlencoded::byte_serialize(encoded_mac.as_bytes()).collect::<String>();

    Ok((headers_list, escaped))
}
