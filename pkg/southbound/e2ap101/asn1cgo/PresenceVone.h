/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-CommonDataTypes"
 * 	found in "../../../../api/e2ap/v1beta2/e2ap-v01.01.00.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_PresenceVone_H_
#define	_PresenceVone_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum PresenceVone {
	PresenceVone_optional	= 0,
	PresenceVone_conditional	= 1,
	PresenceVone_mandatory	= 2
} e_PresenceVone;

/* PresenceVone */
typedef long	 PresenceVone_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_PresenceVone_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_PresenceVone;
extern const asn_INTEGER_specifics_t asn_SPC_PresenceVone_specs_1;
asn_struct_free_f PresenceVone_free;
asn_struct_print_f PresenceVone_print;
asn_constr_check_f PresenceVone_constraint;
ber_type_decoder_f PresenceVone_decode_ber;
der_type_encoder_f PresenceVone_encode_der;
xer_type_decoder_f PresenceVone_decode_xer;
xer_type_encoder_f PresenceVone_encode_xer;
per_type_decoder_f PresenceVone_decode_uper;
per_type_encoder_f PresenceVone_encode_uper;
per_type_decoder_f PresenceVone_decode_aper;
per_type_encoder_f PresenceVone_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _PresenceVone_H_ */
#include "asn_internal.h"
